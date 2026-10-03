#!/bin/bash
# scripts/release.sh - Release Notes Extraction Automation

set -e

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

usage() {
    echo "Usage: $0 [version] [--dry-run]"
    echo "Example: $0 v0.3.0"
    exit 1
}

DRY_RUN=false
VERSION=""

while [[ $# -gt 0 ]]; do
    case $1 in
        --dry-run)
            DRY_RUN=true
            shift
            ;;
        -n)
            DRY_RUN=true
            shift
            ;;
        *)
            if [[ -z "$VERSION" ]]; then
                VERSION=$1
            else
                usage
            fi
            shift
            ;;
    esac
done

if [[ -z "$VERSION" ]]; then
    # Try to determine version from git tags if not provided
    VERSION=$(git describe --tags --abbrev=0 2>/dev/null || echo "")
    if [[ -z "$VERSION" ]]; then
        echo -e "${RED}Error: No version provided and no tags found.${NC}"
        usage
    fi
    echo -e "${YELLOW}Using latest tag: $VERSION${NC}"
fi

CHANGELOG_FILE="CHANGELOG.md"

if [[ ! -f "$CHANGELOG_FILE" ]]; then
    echo -e "${RED}Error: $CHANGELOG_FILE not found.${NC}"
    exit 1
fi

extract_changelog() {
    local ver=$1
    local clean_ver="${ver#v}"
    
    # Find the line where the version header starts (matching either [vX.Y.Z] or [X.Y.Z])
    local start_line=$(grep -nE "^## \[v?${clean_ver}\]" "$CHANGELOG_FILE" | head -n 1 | cut -d: -f1)
    
    if [[ -z "$start_line" ]]; then
        echo -e "${YELLOW}Warning: Version header for '$ver' not found in $CHANGELOG_FILE${NC}" >&2
        echo "## ${ver}"
        echo "No release notes available."
        return 0
    fi

    # Find the next header (## [...) to determine the end of the section
    local next_header_line=$(tail -n +$((start_line + 1)) "$CHANGELOG_FILE" | grep -nE "^## \[" | head -n 1 | cut -d: -f1)
    
    if [[ -z "$next_header_line" ]]; then
        sed -n "$((start_line + 1)),\$p" "$CHANGELOG_FILE"
    else
        local end_relative=$((next_header_line - 1))
        local end_absolute=$((start_line + end_relative))
        sed -n "$((start_line + 1)),$end_absolute p" "$CHANGELOG_FILE"
    fi
}

echo -e "${GREEN}Processing release notes for $VERSION...${NC}"

if $DRY_RUN; then
    echo -e "${YELLOW}*** DRY RUN MODE ***${NC}"
    extract_changelog "$VERSION"
    echo -e "${YELLOW}*** DRY RUN COMPLETE ***${NC}"
else
    extract_changelog "$VERSION" > RELEASE_NOTES.md
    echo -e "${GREEN}Release notes written to RELEASE_NOTES.md${NC}"
fi
